
Write-Output "=== NetAdapter ==="
#Get-NetAdapter | ConvertTo-Json
Get-NetAdapter | Format-Table -AutoSize

Write-Output "=== NetAdapterBinding-WiFi ==="
#Get-NetAdapterBinding
#Get-NetAdapterBinding -Name Wi-Fi | ConvertTo-Json
Get-NetAdapterBinding -Name Wi-Fi | Format-Table -AutoSize

Write-Output "=== NetAdapterBinding-Ethernet ==="
#Get-NetAdapterBinding -Name Ethernet | ConvertTo-Json
Get-NetAdapterBinding -Name Ethernet | Format-Table -AutoSize

#Get-NetAdapterBinding -Name Wi-Fi -ComponentID vms_pp

Write-Output "=== VMSwitch ==="
#Get-VMSwitch | ConvertTo-Json
Get-VMSwitch | Format-Table -AutoSize

Write-Output "=== Process-Wsl ==="
#Get-Process -Name WSL -ErrorAction SilentlyContinue
Get-Process -Name WSL -ErrorAction SilentlyContinue

Write-Output "=== Done ==="

