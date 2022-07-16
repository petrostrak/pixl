#### Pixl
Pixl is a pixel art cross platform desktop app editor that is written in Go and uses the Fyne toolkit.

To use [fyne](https://developer.fyne.io/) library you need to install gcc. On debian a couple of more packages are required. `sudo apt install libgl1-mesa-dev xorg-dev`

##### To install pixl on debian:

* Extract pixl.tar.xz
    
    `mkdir pixl && tar -xvf pixl.tar.xz -C pixl`

* Mount and run make file:

    `cd pixl && sudo make install`