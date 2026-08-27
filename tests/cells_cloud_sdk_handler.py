import compileall
import json




class CellsCloudSDKHandler():
    def __init__(self, logger : logging.Logger):
        self.__ai_client = AIClient(os.getenv("ASPOSE_AI_URI"), os.getenv("ASPOSE_AI_KEY"), os.getenv("ASPOSE_AI_MODEL"))
        self.__logger = logger
        pass

    def auto_commit(self,repo_folder :str, tag_name :str ):
        auto_commit_with_gitpython(repo_folder,tag_name)
        pass
    def tag(self,repo_folder :str,tag_name :str):
        add_tag(repo_folder,tag_name)
        pass
    def push(self,repo_folder :str,branch_name:str="master"):
        push_code(repo_folder,"origin",branch_name)
    def release_with_github(self,product_name:str, version:str, lang:str ,repo_folder :str,repo_name :str,branch_name:str="master"):
        issue_content = self._get_release_issues_content()
        release_content = self.__ai_client.chat( get_user_prompt_with_release_note(version,lang,issue_content))
        auto_publish(repo_folder,repo_name,version,product_name,release_content)
        pass
    def update_readme(self,repo_folder :str, version:str):
        readme_path = Path(repo_folder)/"README.md"
        readme_content = readme_path.read_text()
        readme_release_content = self.__ai_client.chat(get_user_prompt_with_release_version_informat_from_readme(readme_content))
        change_log_path = Path(repo_folder)/"CHANGELOG.md"
        insert_lines_simple(change_log_path,readme_release_content+"\n")
        issue_content = self._get_release_issues_content()
        new_readme_content = self.__ai_client.chat(get_user_prompt_with_update_readme(version,issue_content,readme_content))
        write_file(readme_path,new_readme_content)
        pass
    def start_test(self, lang:str,work_folder):

        # for language in ["go", "java", "net", "node", "perl","php","python","ruby"]:
        #     method_name = f"_start_sdk_test_with_{language}"
        #     if hasattr(self, method_name):
        #         method = getattr(self, method_name)
        #         method(work_folder)
        # pass
        method_name = f"_start_sdk_test_with_{lang}"
        if hasattr(self, method_name):
            method = getattr(self, method_name)
            return method(work_folder)
    def _get_test_group_names(self, configuration_folder):
        group_names = []
        for file_path in Path(configuration_folder).glob("*.json"):
            if file_path.is_file():
                json_data = load_config(str(file_path))
                group_names.append(json_data["Name"])
        return group_names

    def _start_sdk_test_with_go(self , work_folder:str = ""):
        work_path = Path(work_folder)
        test_work_folder = str( work_path/"integrationtests")
        test_script_file_path = str( work_path/"integrationtests/test_scripts.ps1")

        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Go package.")
        build_cmd = ["go", "build", "."]
        result = run_command(build_cmd, work_folder)
        if result["success"] :
            self.__logger.info("Beginning tidy test case for Aspose.Cells Cloud SDK for Go. ")
            mod_cmds = ["go", "mod", "tidy"]
            result = run_command(mod_cmds, test_work_folder)
            if result["success"] :
                self.__logger.info("Beginning fetch script file about Aspose.Cells Cloud SDK for Go Integration Tests.")
                test_cmds = []
                with open(test_script_file_path, 'r',
                          encoding='utf-8') as f:
                    for line in f:
                        line = line.rstrip('\n')
                        cmd = line.split()
                        test_cmds.append(cmd)
                error_info = "Error Information:\n"
                std_out_info = ""
                self.__logger.info("Beginning Aspose.Cells Cloud SDK for Go Integration Tests.")
                for test_cmd in test_cmds:
                    self.__logger.info(test_cmd)
                    result = run_command(test_cmd,test_work_folder)
                    std_out_info += result["stdout"]
                    if not result["success"]:
                        error_info += result["stderr"]

                response = self.__ai_client.chat(
                    get_user_prompt_with_sdk_integration_test(std_out_info + error_info))
                response_json = json.loads( response)
                if response_json["total_tests"] != response_json["passed"]:
                    redmine = get_redmine()
                    redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Go integrate test fail. ",
                                      description=json.dumps(response_json["failures"]))
                    return False
                else:
                    return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat( get_user_prompt_with_build_project(result["stderr"]))
        redmine = get_redmine()
        redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Go build fail.",
                          description= json.dumps(response))
        return False

    def _start_sdk_test_with_java(self, work_folder:str ="",test_data_configuration_folder:str =""):
        work_path = Path(work_folder)
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Go package.")
        build_cmd = ["mvn", "clean", "compile"]
        result = run_command(build_cmd, work_folder)
        if result["success"]:
            build_cmd = ["mvn", "clean", "test-compile"]
            result = run_command(build_cmd, work_folder)
            if result["success"]:
                group_names = self._get_test_group_names(test_data_configuration_folder)
                error_info = "Error Information:\n"
                std_out_info = ""
                for group_name in group_names:
                    test_cmd = ["mvn", "clean","test", f"-Dgroups={group_name}"]
                    result = run_command(test_cmd,work_folder)
                    std_out_info += result["stdout"]
                    if not result["success"]:
                        error_info += result["stderr"]
                response =json.loads( self.__ai_client.chat(
                    get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
                if response["total_tests"] != response["passed"]:
                    redmine = get_redmine()
                    redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Java integrate test fail. ",
                                      description=json.dumps(response["failures"]))
                    return False
                else:
                    return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat(get_user_prompt_with_build_project(result["stderr"]))
        redmine = get_redmine()
        redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Java build fail.",
                          description=json.dumps(response))
        return False

    def _start_sdk_test_with_net(self, work_folder:str ="",test_data_configuration_folder:str =""):
        work_path = Path(work_folder)
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Net package.")
        build_cmd = ["dotnet", "build", ".\Aspose.Cells.Cloud.SDK\Aspose.Cells.Cloud.SDK.csproj"]
        result = run_command(build_cmd, work_folder)
        if result["success"]:
            build_cmd = ["dotnet", "build", ".\Aspose.Cells.Cloud.SDK.Test\Aspose.Cells.Cloud.SDK.Test.csproj"]
            result = run_command(build_cmd, work_folder)
            if result["success"]:
                group_names = self._get_test_group_names(test_data_configuration_folder)
                error_info = "Error Information:\n"
                std_out_info = ""
                for group_name in group_names:
                    test_cmd = ["dotnet", "test",".\Aspose.Cells.Cloud.SDK.Test\Aspose.Cells.Cloud.SDK.Test.csproj", "--filter",f"TestCategory={group_name}"]
                    result = run_command(test_cmd,work_folder)
                    std_out_info += result["stdout"]
                    if not result["success"]:
                        error_info += result["stderr"]
                response = json.loads( self.__ai_client.chat(
                    get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
                if response["total_tests"] != response["passed"]:
                    redmine = get_redmine()
                    redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Net integrate test fail. ",
                                      description=json.dumps(response["failures"]))
                    return False
                else:
                    return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat(get_user_prompt_with_build_project(result["stderr"]))
        redmine = get_redmine()
        redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Net build fail.",
                          description=json.dumps(response))
        return False

    def _start_sdk_test_with_node(self, work_folder:str ="",test_data_configuration_folder:str =""):
        work_path = Path(work_folder)
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Net package.")
        build_cmd = ["dotnet", "build", ".\Aspose.Cells.Cloud.SDK\Aspose.Cells.Cloud.SDK.csproj"]
        result = run_command(build_cmd, work_folder)
        if result["success"]:
            build_cmd = ["dotnet", "build", ".\Aspose.Cells.Cloud.SDK.Test\Aspose.Cells.Cloud.SDK.Test.csproj"]
            result = run_command(build_cmd, work_folder)
            if result["success"]:
                group_names = self._get_test_group_names(test_data_configuration_folder)
                error_info = "Error Information:\n"
                std_out_info = ""
                for group_name in group_names:
                    test_cmd = ["dotnet", "test",".\Aspose.Cells.Cloud.SDK.Test\Aspose.Cells.Cloud.SDK.Test.csproj", "--filter",f"TestCategory={group_name}"]
                    result = run_command(test_cmd,work_folder)
                    std_out_info += result["stdout"]
                    if not result["success"]:
                        error_info += result["stderr"]
                response =json.loads( self.__ai_client.chat(
                    get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
                if response["total_tests"] != response["passed"]:
                    redmine = get_redmine()
                    redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Java integrate test fail. ",
                                      description=json.dumps(response["failures"]))
                    return False
                else:
                    return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat(get_user_prompt_with_build_project(result["stderr"]))
        redmine = get_redmine()
        redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Java build fail.",
                          description=json.dumps(response))
        return False

    def _start_sdk_test_with_perl(self, work_folder:str =""):
        work_path = Path(work_folder)
        test_case_folder = work_path / "t"
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Perl package.")
        error_info = "Error Information:\n"
        std_out_info = ""
        for test_file in test_case_folder.glob("*.pl"):
            test_cmd = ["prove", "-t", str(test_file)]
            result = run_command(test_cmd,work_folder)
            std_out_info += result["stdout"]
            if not result["success"]:
                error_info += result["stderr"]
        response = json.loads(self.__ai_client.chat(
            get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
        if response["total_tests"] != response["passed"]:
            redmine = get_redmine()
            redmine.new_issue("13", "Fix few bugs of the Aspose.Cells Cloud SDK for Perl integrate test fail. ",
                              description=json.dumps(response["failures"]))
            return False
        else:
            return True
        return False

    def _start_sdk_test_with_php(self, work_folder:str =""):
        work_path = Path(work_folder)
        test_case_folder = work_path/"integrationtests"
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for PHP package.")

        # build_cmd = ["Get-ChildItem", "lib", "-Recurse", "-Filter", "*.php", "|", "ForEach-Object", "{", "php", "-l", "$_.FullName", "}"]
        build_cmd = ["composer.bat", "install", "--ignore-platform-reqs"]
        result = run_command(build_cmd, work_folder)
        if result["success"]:
            self.__logger.info( "Run Command: " +  " ".join(build_cmd))
            error_info = "Error Information:\n"
            std_out_info = ""
            for test_file_path in test_case_folder.rglob("*.php"):
                print(test_file_path)
                if test_file_path.name.endswith("CellsApiTestBase.php"):
                    continue
                test_cmd = ["vendor\\bin\\phpunit.bat", str(test_file_path)]
                self.__logger.info(f"Work folder: {work_folder}.")
                self.__logger.info("Running: " + " ".join(test_cmd))
                result = run_command(test_cmd,work_folder)
                std_out_info += result["stdout"]
                if not result["success"]:
                    error_info += result["stderr"]
                    self.__logger.info("Run Fail: " + result["stderr"])
            response = json.loads(self.__ai_client.chat(
                get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
            if response["total_tests"] != response["passed"]:
                # new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for PHP integrate test fail. ",
                #                   description=json.dumps(response["failures"]))
                return False
            else:
                return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat(get_user_prompt_with_build_project(result["stderr"]))
        new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for PHP build fail.",
                          description=json.dumps(response))
        return False

    def _start_sdk_test_with_python(self, work_folder:str ="",test_data_configuration_folder:str=""):
        work_path = Path(work_folder)
        test_case_folder = work_path/"integrationtests"
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Python package.")
        success = compileall.compile_dir("asposecellscloud", quiet=3)
        if success:
            self.__logger.info("Build Aspose.Cells Cloud SDK for Python successes.")
            error_info = "Error Information:\n"
            std_out_info = ""
            # group_names = self._get_test_group_names(test_data_configuration_folder)
            # for group_name in group_names:
            # self.__logger.info( group_name)
            # test_cmd = ["python",f"{test_case_folder}\tests_{to_snake_case(group_name)}.py" ]
            for test_file_path in test_case_folder.glob("tests_*.py"):
                test_cmd = ["python", test_file_path]
                self.__logger.info(f"Run test case file: {test_file_path}")
                result = run_command(test_cmd,work_folder)
                std_out_info += result["stdout"]
                if not result["success"]:
                    error_info += result["stderr"]
                    self.__logger.info( result["stderr"])
            response = json.loads( self.__ai_client.chat(
                get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
            self.__logger.info( json.dumps(response))
            if response["total_tests"] != response["passed"]:
                new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for Python integrate test fail. ",
                                  description=json.dumps(response["failures"]))
                return False
            else:
                return True
        self.__logger.info("There is a grammatical error")
        new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for Python build fail.",
                          description="There is a grammatical error")
        return False

    def _start_sdk_test_with_ruby(self, work_folder:str =""):
        work_path = Path(work_folder)
        test_case_folder = work_path/"spec"
        self.__logger.info("Beginning build Aspose.Cells Cloud SDK for Ruby package.")

        build_pwsh = " Get-ChildItem lib -Recurse -Filter *.rb | ForEach-Object { ruby -c $_.FullName }"
        result = run_command(build_pwsh, work_folder)
        if result["success"]:
            self.__logger.info("Build Aspose.Cells Cloud SDK for Ruby package succeeded.")
            error_info = "Error Information:\n"
            std_out_info = ""
            for test_file_path in test_case_folder.rglob("*.rb"):
                if test_file_path.name.endswith("spec_helper.rb"):
                    continue
                self.__logger.info(f"Run test case file: {test_file_path}")
                test_cmd = ["rspec", str(test_file_path)]
                result = run_command(test_cmd,work_folder)
                std_out_info += result["stdout"]
                if not result["success"]:
                    error_info += result["stderr"]
                    self.__logger.info(result["stderr"])
            response = json.loads( self.__ai_client.chat(
                get_user_prompt_with_sdk_integration_test(std_out_info + error_info)))
            if response["total_tests"] != response["passed"]:
                new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for Ruby integrate test fail. ",
                                  description=json.dumps(response["failures"]))
                return False
            else:
                return True
        self.__logger.info(result["stderr"])
        response = self.__ai_client.chat(get_user_prompt_with_build_project(result["stderr"]))
        new_issue(get_redmine(),"13", "Fix few bugs of the Aspose.Cells Cloud SDK for Ruby build fail.",
                          description=json.dumps(response))
        return False


    def _get_release_issues_content(self):
        issues = get_project_release_issues(get_redmine(),"13")
        issue_content ="| **Summary**                                                                                                   | **Category** |\n"
        issue_content = issue_content + "| :------------------------------------------------------------------------------------------------------------ | :----------- |"
        for issue in issues:
            issue_content += format_issue_data(issue)
        return issue_content