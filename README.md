# Virsh Device Daemon

This was created to allow utilizing one keyboard and mouse between a host and guest with KVM/QEMU. It is intended to
be run as a server on the host which the guest communicates with via a REST API. The guest must have network access
to the host. Both the client and server are contained in the same binary.

### Usage

Build the Linux and Windows binaries:
```
go get github.com/fapiko/virsh-device-daemon
cd $GOPATH/src/github.com/fapiko/virsh-device-daemon
govendor sync
make build
```

Create an xml file for each device you wish to attach & detach from the VM. The command ```lsusb``` can be used to 
find the vendor and product IDs for the USB devices to manage.

For example I have one file each for my keyboard and mouse (though this should work with any USB device). Here is my
keyboard file:
```
<hostdev mode='subsystem' type='usb' managed='yes'>
   <source>
       <vendor id='0x046d'/>
       <product id='0xc22d'/>
   </source>
</hostdev>
```

Next, start the server on the host machine:
```./virsh-device-daemon -s -f device-files/keyboard.xml -f device-files/mouse.xml -n win8.1```

  * The ```-s``` flag tells it to start in server mode
  * The ```-f``` flags indicate the location of the device files to use
  * The ```-n``` flag gives it the name of the VM to attach/detach devices to
  
Now setup a global hotkey on the host machine. In Ubuntu this can be done by opening the ```Keyboard``` app and
clicking on the shortcuts tab. On the host I want to attach my keyboard and mouse to the VM, so I add the command:
```/opt/gopath/src/github.com/fapiko/virsh-device-daemon/virsh-device-daemon -a```

Next copy the binary to the Guest VM. Since my guest is Windows, I use shortcut placed on the desktop for the global
hotkey that calls a batch script which calls the virsh-device-daemon-win64.exe binary. The batch script is as
follows:
```/opt/gopath/src/github.com/fapiko/virsh-device-daemon/virsh-device-daemon -h 192.168.1.100:7654 -d```

  * The ```-h``` flag indicates the hostname of the service running on the host machine
  * The ```-d``` flag tells it to detach the USB devices and give control back to the guest
  
That's it! Now you can use your shortcut key on the host machine when you want to attach your devices to the guest,
and vice versa.