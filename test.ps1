Write-Host "Start Aspose.Cells Cloud SDK for GO"
    $StartTime = Get-Date
    
    [string[]]$lines = go test -v  --timeout 50m
    $passed = 0
    $total =0
    foreach( $line in  $lines)
    {
        if(($line -match  "=== RUN"))
        {
            $total++
        }      
        if(($line -match  "--- PASS") ) 
        {
            $passed++
        }
    }
    $failed = $total - $passed
    $EndTime = Get-Date
    $timespan ="{0:N2}" -f (New-TimeSpan $StartTime  $EndTime).TotalSeconds
    Write-Host "Spent ${timespan}s on finishing test. Result : Total ${total}, Passed ${passed} , Failed ${failed} ."
    