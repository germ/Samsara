First off, thanks for taking interest in Samsara. This means a lot to me!

I’ve attached a tarball containing Samsara and all components needed for compilation. I’ve also included a related project TransmissionLimiting, because running this 24/7 with housemates leads to a ton of hard shutdowns :)
So the code really isn’t organized that well, most of the comments (if many) are in the code, a ton of it is hacks. Feel free to email me with any questions and what not and I’ll clear it up, and apologies again for the quality.

The way I have it set up Samsara is run every 6 hours. The flow kind of looks like this. 

[ (Program : Trigger) Action ]
(Samsara: Six hour timer) LastFM suggestion fetching, rottenTomato fetching, adding torrents to Transmission
(Transmission: Torrent Finish) move files to iTunes autoadd folder
(iTunes: New Songs added) Update smart playlists
(iTunes: iDevice on Wifi) AutoSync

Dirs:
iTunesDB: 				Lib for extracting iTunes Library info (incomplete)
libTransmission:			Basic torrent management
rottenTomatoes:			Wrapper for the rottenTomatoes api
Samsara:					Zero-Interaction downloader
TransmissionLimiting:		Web interface for keeping roommates sane

Any reference to Germ should be replaced with the current user. Theres a nice smattering of these hard coded refs kicking around. Sorry!

In order for this to really work, everything needs to be run as a daemon. I’ve included a sampling of plists I was using to get this working. The TransmissionLimiting web server was run as root, Samsara/Trans as local user. Another idea I tried was to run iTunes in the same way, you could modify Trans.plist if needed.

TransmissionLimiting just needs to be bolted onto net/http and daemonized, homeSkynet.plist has what I was using to control this. 
MoveToiTunes.sh is called after Transmission downloads a torrent by Transmission. This happens before seeding though. If you can, try and find a way to allow proper seeding. The last thing I want is this thing destroying swarm health even more.

I’ve added a example of one of my smart playlists for syncing my device. A good idea is to add rules to exclude tracks from the other playlists for the device. Memory remembers this being a snag. 

Please Note:
This entire project hinges on LastFM, it needs to be installed, configured and actively used before it becomes useful. I would recommend at least a day or two of listening before allowing Samsara to download preferred music. I know there is a way to find the total scribbles via the API. If it is under some number, just have it download the mainstream music. Although once it is conditioned you can delete your entire ~/.samsaraDB and get an entirely new library (just modify con.yaml to your liking). Cmd+Shift+A, Cmd+Shift+Opt+Delete and welcome to a new library!

Finding a way to install this is just entirely beyond my scope. It will need some kind of script or pkg installer though.
My advice, just read the code a few times to get a hang of it. Then dive it and do your freaky thing.

Have a great day, and thank you again!

