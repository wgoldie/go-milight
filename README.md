go-milight
==========

This is a project to experiment with and control Milight/LimitlessLED wifi lightbulbs. 
These are kinda like the philips hue but only support 256 colors (yay byte overflow) and cost less than $20 each instead of $100s. 
Also they have much less useful/acccurate/lucid documentation than I imagine Philips or similar bulbs have,
and probably are way worse from a quality of light perspective.
But I'm a broke student so they're perfect.


I plan to create an animation/tweening framework and a couple visualizers for music, video, etc. I may also do a Go package for general use of the bulbs in code.
Some people have done similar visualizers but in my experience they haven't included some features that are imo vital, such as:
-	Lag compensation for local or delayed audio and video tracks
-	Fine control over color ranges allowed by the user
-	Multiple sending of control messages to compensate for dropped packets or whatever it is that makes bulbs sometimes miss color changes.


I have no idea if any of this will be compatible with other users' setups. 