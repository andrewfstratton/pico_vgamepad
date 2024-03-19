# pico_vgamepad
Rpi Pico Virtual Gamepad controlled through serial commands - for Quando

Built with TinyGo (currently) this makes the PRpi Pico act as a gamepad that can be controlled by(simple) serial commands.

## Limitations

- Only tested as DirectInput for Windows
- Steam can be used to map the inputs to XInput so can be used with most games this way.
  - Note that Xinput is not included directly

## Commands

There are currently only a few commands:

- b99 - will release button 99, where 99 is typically in the range 1 to 16
- B99 - will press button 99
- [X|Y|Z|x|y|z]99999 - will change the X/Y/Z/x/y/z axis value to 99999, which has a range of -32767 to 32767 where 0 is the middle position
  - _The triggers will still use the whole range, so -32767 means 'at rest'_
