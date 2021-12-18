# Notes

## Screen Saver (Blanking) on Pi

```shell
# To see the X settings:
xset q

# To change the blanking timeout (blank after an hour):
xset s 3600 3600
xset dpms 3600 3600 3600
```

## Putting Keys on Servers

```shell
# To see which keys:
aws s3 ls s3://briancsparks-files/keys/

# Get a specific key into the authorized_keys file
curl -L https://briancsparks-files.s3.amazonaws.com/keys/id_rsa.pub >> ${HOME}/.ssh/authorized_keys
```

## Dev Server

In order to use CLion remote, must setup the server for development. The only
thing that was difficultish was to install CMake 3.20, the version in the repos
was something earlier.

General dev-server setup (for C/C++):

```shell
#sudo apt-get install -y cmake gcc gdb tree curl git htop libssl-dev build-essential
sudo apt-get install -y gcc gdb tree curl git htop libssl-dev build-essential
```

Upgrading to CMake 3.20

```shell
#cmake -version
#sudo apt-get remove -y cmake
#sudo apt autoremove

curl -L 'https://github.com/Kitware/CMake/releases/download/v3.20.0/cmake-3.20.0.tar.gz' | tar -zxv
cd cmake-3.20.0/
./bootstrap
make
sudo make install
```








