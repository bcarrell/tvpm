tvpm
====

tvpm is a television episode package manager for newsgroups.

The goal of the project is to provide a command line tool for managing your media, 
inspired by all of the great package managers out there.

**This tool should be considered alpha and has the potential to break or change 
massively.**

This tool depends on [Sabnzbd](http://sabnzbd.org/).

Quickstart
==========

	git clone git@github.com:bcarrell/tvpm.git $GOPATH/src/github.com/bcarrell/tvpm
	cd $GOPATH/src/github.com/bcarrell/tvpm && go install

	echo 'export TRAKT_API_KEY="your_api_key_here"' >> $HOME/.zshrc
	echo 'export TVPM_DB_PATH="your_path_here"' >> $HOME/.zshrc
	echo 'export SABNZBD_URL="your_running_sabnzbd_url_here"' >> $HOME/.zshrc
	echo 'export SABNZBD_API_KEY="your_api_key_here"' >> $HOME/.zshrc

	tvpm


For tvpm to work properly, set the following env variables for bash/zsh/whatever:

* **TRAKT_API_KEY** -- your API key for trakt.tv
* **TVPM_DB_PATH** -- absolute path where you want the tvpm sqlite3 db to be stored
* **SABNZBD_URL** -- the url of your running sabnzbd installation
* **SABNZBD_API_KEY** -- your sabnzbd api key


Available commands
------------------
All commands are available via `tvpm --help` or just `tvpm`.

	tvpm add-indexer <indexer url> --apiKey=<indexer api key>
The above command adds an available newsgroup indexer to your tvpm database.

	tvpm find-series <series>
The above command uses trakt to find a series.  Currently, this command is somewhat of 
a dead end.

	tvpm find <query>
The above command searches all of your indexers for the specified tv episode.  Format 
your search like `tvpm find game+of+thrones` or `tvpm find game-of-thrones`.  Underscores 
should work too.  Then, it'll give you the option to send the one you want to sabnzbd.



Contributing
------------

I welcome any and all pull requests; I'll try to keep issues created for things 
I find important, but if you have suggestions, feel free.