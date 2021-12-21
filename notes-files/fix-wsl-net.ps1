
# distro            (Ubuntu-20.04, Ubuntu-18.04, Alpine)
# linux-username    (sparksb, scotty)
# AdapterName       (Ethernet, Wi-Fi)
# subnet CIDR
# upstream

$adapter = "Ethernet"
#$adapter = "Wi-Fi"
$distro = "Ubuntu-20.04"
$username = "sparksb"
$cidr = "192.168.86.240/24"
$upstream = "192.168.86.1"

Write-Output "Starting WSL, without shell" ;

# wsl -d Ubuntu-20.04 -u sparksb --exec exit
wsl -d $distro -u $username --exec exit

# Wait for WSL switch to come up
$started = $false
Do {
  $status = Get-VMSwitch WSL -ErrorAction SilentlyContinue ;
  if (!($status)) {
    Write-Output "Waiting for WSL switch" ;
    Start-Sleep 1 ;
  }
  Else {
    $started = $true;

  }
}
Until ( $started )

Write-Output "WSL Switch is up, bridging." ;

# Set-NetAdapterBinding -Name Wi-Fi -ComponentID vms_pp -Enabled $False ;
# Set-VMSwitch WSL -NetAdapterName Wi-Fi ;

# Set-NetAdapterBinding -Name Ethernet -ComponentID vms_pp -Enabled $False ;
# Set-VMSwitch WSL -NetAdapterName Ethernet ;

Set-NetAdapterBinding -Name $adapter -ComponentID vms_pp -Enabled $False ;
Set-VMSwitch WSL -NetAdapterName $adapter ;

         # TODO? This is how you fixup VMs
#        # Hook all Hyper V VMs to WSL network => avoid network performance issues.
#        Write-Output  "Getting all Hyper V machines to use WSL Switch" >> $logPath ;
#        Get-VM | Get-VMNetworkAdapter | Connect-VMNetworkAdapter -SwitchName "WSL" ;

# configure-wsl2-net.sh 192.168.86.240/24 192.168.86.1
# wsl -d Ubuntu-20.04 -u sparksb /home/sparksb/configure-wsl2-net.sh 192.168.86.240/24 192.168.86.1
# Write-Output Start-Process -FilePath "wsl.exe" -ArgumentList "-d Ubuntu-20.04 -u sparksb /home/sparksb/configure-wsl2-net.sh 192.168.86.240/24 192.168.86.1"
Write-Output Start-Process -FilePath "wsl.exe" -ArgumentList "-d $distro -u $username /home/$username/configure-wsl2-net.sh $cidr $upstream"


# # ---------------------------------------------------------------------
# $ cat configure-wsl2-net.sh
# # ---------------------------------------------------------------------
# #!/bin/bash
# set -x
#
# CIDR="192.168.86.239/24"
# UPSTREAM="192.168.86.1"
#
# [[ -n $1 ]] && CIDR="$1"
# [[ -n $2 ]] && UPSTREAM="$2"
#
# echo "booya | $CIDR |   | $UPSTREAMPSTREAM |" > /home/sparksb/log.txt
#
# # 192.168.86.90 your WSL2 fixed address outside of DHCP / 192.168.86.1 your router address
# sudo ip addr flush eth0
# #sudo ip addr add 192.168.86.240/24 brd + dev eth0
# sudo ip addr add ${CIDR} brd + dev eth0
# sudo ip route delete default
# #sudo ip route add default via 192.168.86.1
# sudo ip route add default via "${UPSTREAM}"
#
# # run docker deamon and enjoy your hassle free containers.
# #sudo dockerd
# # ---------------------------------------------------------------------




