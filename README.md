tvpm
====

tvpm is a television episode package manager for newsgroups.

The goal of the project is to provide a command line tool for managing your media.


Environment variables
---------------------

For tvpm to work properly, set the following env variables for bash/zsh/whatever:

* TRAKT_API_KEY -- your API key for trakt.tv
* TVPM_DB_PATH -- absolute path where you want the tvpm sqlite3 db to be stored
* SABNZBD_URL -- the url of your running sabnzbd installation
* SABNZBD_API_KEY -- your sabnzbd api key