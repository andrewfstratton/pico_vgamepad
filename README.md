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

To be implemented - Axis:
- [X|Y|Z]99999 - will change the X/Y/Z axis value to 99999, which has a range of 1 to 65535 where 32768 is the middle position
- Also x|y|z are the 'rotation' axis
- N.B. The triggers will still use the whole range, so 1 means 'at rest'

~~May decide to use 1 to 255 only for triggers~~
