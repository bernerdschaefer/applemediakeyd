applemediakeyd
--------------

applemediakeyd is a daemon for linux,
which listens for media key events
and does the right thing.

## Backlight brightness

applemediakeyd reads and writes
to the backlight brightness file
in the sysfs.

## Keyboard backlight

applemediakeyd reads and writes
to the keyboard backlight brightness file
in the sysfs.

## Volume

applemediakeyd uses the OSS API
to maintain volume.

The OSS API is actually an emulation layer
which maps to ALSA driver commands.
But it allows us to use native Go code
with syscalls for `ioctl`,
instead of requiring `cgo` and `alsa-lib`.

## TODO

  - [ ] extract brightness into package
  - [ ] connect previous, play/pause, next buttons to [MPRIS]
  - [ ] adjust based on ambient sensor.

  [MPRIS]: http://specifications.freedesktop.org/mpris-spec/latest/

## License

sgheme is Copyright (c) 2015 Bernerd Schaefer.
It is free software, and may be redistributed
under the terms specified in the [LICENSE] file.
