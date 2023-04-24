$date=$(Get-Date -Format yyyyMMddHHmmss)
Write-Output $date
if (-not (Test-Path TestResult)){
    New-Item -Path "TestResult" -ItemType  "directory"
}

go test -v --timeout 50m | go-junit-report | Out-File -FilePath ./TestResult/IntegrationTestResult$date.xml
