# 这个版本做了哪些修改

1. 增加了专辑恢复功能(只有web下生效).
   因为我用navidrome听MOOC, 原来的版本, 如果你播放专辑, 就会从第一个音频开始播放, 我希望能重上一次最后播放的那一个音频开始.
2. 增加了自定义专辑名, 专辑艺术家的功能.
   我下载了U选1000, 有1000首歌, 如果用原来的版本, 就会分成非常多的专辑, 官方给的解决方案是修改TAG <https://www.navidrome.org/docs/faq/#-i-have-an-album-with-tracks-by-different-artists-why-is-it-broken-up-into-lots-of-separate-albums-each-with-their-own-artist>, 把艺术家名统一修改成"群星".
   但是这样在播放的时候就看不到具体的演唱者了, 于是添加了这个功能.
   方法是在目录里面放一个album.json文件, 内容如下, 这样这些歌都会放在同一个专辑里, 而且播放的时候演唱者的名字也是原来的.
```json
{ "name": "U选1000", "artist": "群星" }
```

What is different from original version:

1. Added album resume function(web player only).
   Because I use navidrome to listen to MOOC, the original version, if you play the album, it will start from the first audio, I would like to start from the last audio that was played last time.
2. add the function of customising album name, album artist.
   I have downloaded U1000, there are 1000 songs, if I use the original version, it will be divided into many albums, the official solution is to change the TAG <https://www.navidrome.org/docs/faq/#-i-have-an-album-with-tracks-by-different-artists-why-is-it-broken-up-into-lots-of-separate-albums-each-with-their-own-artist>, change the artist name to "Various Artists". However, the artist could not be seen during playback, so this feature was added.
   The way to do this is to put an album.json file in the directory, with the following content, so that all these songs will be put in the same album, and the artist's name will be the same when playing.
```json
{ "name": "U1000", "artist": "Various Artists" }
```

# Build static app

```bash
cp libtag.a libtag_c.a /usr/lib/x86_64-linux-gnu #files come from docker deluan/ci-goreleaser
make buildstatic
```

<a href="https://www.navidrome.org"><img src="resources/logo-192x192.png" alt="Navidrome logo" title="navidrome" align="right" height="60px" /></a>

# Navidrome Music Server &nbsp;[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Tired%20of%20paying%20for%20music%20subscriptions%2C%20and%20not%20finding%20what%20you%20really%20like%3F%20Roll%20your%20own%20streaming%20service%21&url=https://navidrome.org&via=navidrome)

[![Last Release](https://img.shields.io/github/v/release/navidrome/navidrome?logo=github&label=latest&style=flat-square)](https://github.com/navidrome/navidrome/releases)
[![Build](https://img.shields.io/github/actions/workflow/status/navidrome/navidrome/pipeline.yml?branch=master&logo=github&style=flat-square)](https://nightly.link/navidrome/navidrome/workflows/pipeline/master)
[![Downloads](https://img.shields.io/github/downloads/navidrome/navidrome/total?logo=github&style=flat-square)](https://github.com/navidrome/navidrome/releases/latest)
[![Docker Pulls](https://img.shields.io/docker/pulls/deluan/navidrome?logo=docker&label=pulls&style=flat-square)](https://hub.docker.com/r/deluan/navidrome)
[![Dev Chat](https://img.shields.io/discord/671335427726114836?logo=discord&label=discord&style=flat-square)](https://discord.gg/xh7j7yF)
[![Subreddit](https://img.shields.io/reddit/subreddit-subscribers/navidrome?logo=reddit&label=/r/navidrome&style=flat-square)](https://www.reddit.com/r/navidrome/)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0-ff69b4.svg?style=flat-square)](CODE_OF_CONDUCT.md)

Navidrome is an open source web-based music collection server and streamer. It gives you freedom to listen to your
music collection from any browser or mobile device. It's like your personal Spotify!


**Note**: The `master` branch may be in an unstable or even broken state during development. 
Please use [releases](https://github.com/navidrome/navidrome/releases) instead of 
the `master` branch in order to get a stable set of binaries.

## [Check out our Live Demo!](https://www.navidrome.org/demo/)

__Any feedback is welcome!__ If you need/want a new feature, find a bug or think of any way to improve Navidrome, 
please file a [GitHub issue](https://github.com/navidrome/navidrome/issues) or join the discussion in our 
[Subreddit](https://www.reddit.com/r/navidrome/). If you want to contribute to the project in any other way 
([ui/backend dev](https://www.navidrome.org/docs/developers/), 
[translations](https://www.navidrome.org/docs/developers/translations/), 
[themes](https://www.navidrome.org/docs/developers/creating-themes)), please join the chat in our 
[Discord server](https://discord.gg/xh7j7yF). 

## Installation

See instructions on the [project's website](https://www.navidrome.org/docs/installation/)

## Cloud Hosting

[PikaPods](https://www.pikapods.com) has partnered with us to offer you an 
[officially supported, cloud-hosted solution](https://www.navidrome.org/docs/installation/managed/#pikapods). 
A share of the revenue helps fund the development of Navidrome at no additional cost for you.

[![PikaPods](https://www.pikapods.com/static/run-button.svg)](https://www.pikapods.com/pods?run=navidrome)

## Features
 
 - Handles very **large music collections**
 - Streams virtually **any audio format** available
 - Reads and uses all your beautifully curated **metadata**
 - Great support for **compilations** (Various Artists albums) and **box sets** (multi-disc albums)
 - **Multi-user**, each user has their own play counts, playlists, favourites, etc...
 - Very **low resource usage**
 - **Multi-platform**, runs on macOS, Linux and Windows. **Docker** images are also provided
 - Ready to use binaries for all major platforms, including **Raspberry Pi**
 - Automatically **monitors your library** for changes, importing new files and reloading new metadata 
 - **Themeable**, modern and responsive **Web interface** based on [Material UI](https://material-ui.com)
 - **Compatible** with all Subsonic/Madsonic/Airsonic [clients](https://www.navidrome.org/docs/overview/#apps)
 - **Transcoding** on the fly. Can be set per user/player. **Opus encoding is supported**
 - Translated to **various languages**

## Documentation
All documentation can be found in the project's website: https://www.navidrome.org/docs. 
Here are some useful direct links:

- [Overview](https://www.navidrome.org/docs/overview/)
- [Installation](https://www.navidrome.org/docs/installation/)
  - [Docker](https://www.navidrome.org/docs/installation/docker/)
  - [Binaries](https://www.navidrome.org/docs/installation/pre-built-binaries/)
  - [Build from source](https://www.navidrome.org/docs/installation/build-from-source/)
- [Development](https://www.navidrome.org/docs/developers/)
- [Subsonic API Compatibility](https://www.navidrome.org/docs/developers/subsonic-api/)

## Screenshots

<p align="left">
    <img height="550" src="https://raw.githubusercontent.com/navidrome/navidrome/master/.github/screenshots/ss-mobile-login.png">
    <img height="550" src="https://raw.githubusercontent.com/navidrome/navidrome/master/.github/screenshots/ss-mobile-player.png">
    <img height="550" src="https://raw.githubusercontent.com/navidrome/navidrome/master/.github/screenshots/ss-mobile-album-view.png">
    <img width="550" src="https://raw.githubusercontent.com/navidrome/navidrome/master/.github/screenshots/ss-desktop-player.png">
</p>
