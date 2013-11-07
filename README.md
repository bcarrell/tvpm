tvpm
====

tvpm is a television episode package manager for newsgroups.

The goal of the project is to provide a command line tool for managing your media, 
inspired by all of the great package managers out there.

**This tool should be considered alpha and has the potential to break or change 
massively.**

Available commands
------------------
All commands are available via `tvpm --help` or just `tvpm`.

	tvpm add-indexer <indexer url> --apiKey=<indexer api key>
This command adds an available newsgroup indexer to your tvpm database.

	tvpm find-series <series>
This command uses trakt to find a series.  Currently, this command is somewhat of 
a dead end.

	tvpm find <query>
This command searches all of your indexers for the specified tv episode.  Format 
your search like `tvpm find game+of+thrones` or `tvpm find game-of-thrones`.  Underscores 
should work too.


Environment variables
---------------------

For tvpm to work properly, set the following env variables for bash/zsh/whatever:

* **TRAKT_API_KEY** -- your API key for trakt.tv
* **TVPM_DB_PATH** -- absolute path where you want the tvpm sqlite3 db to be stored
* **SABNZBD_URL** -- the url of your running sabnzbd installation
* **SABNZBD_API_KEY** -- your sabnzbd api key


Building the binary
-------------------

Make sure you have Go installed, and `go install` the repo.