==================
golang-on-dotcloud
==================

Do you have a Go application that you want to run on dotCloud? If so, you have come to the right place. dotCloud doesn't support Go as a native service, but with this custom service recipe you can still use dotCloud for all of your Go needs.

If we get enough people to use this custom service, and show support for it, we might be able to convince dotCloud to add native support for Go. So let us know if you use it, and if you want dotCloud to offer a native Go service. 

Status
======
This is still very early beta so use with caution, and let us know if you have any issues.

Version
=======

Go v1.1.1

How to use
==========
1. Clone this repo::
    
    $  git clone git://github.com/kencochrane/golang-on-dotcloud.git

2. Put your Go source code under the ``src`` directory.

3. Change the ``dotcloud.yml``
    a. Change the ``build_package`` config variable, to the name of your Go package
    b. Change the ``processes`` so that you can run the correct processes
    c. Add or remove ``ports`` depending on your application needs.

4. Create a dotCloud application::
    
    $ dotcloud create -f sandbox mygoapp

5. Push your code to your new dotCloud application::

    $ dotcloud push mygoapp .

6. Enjoy!

What to watch out for
=====================
- Make sure you don't push any of your files from your ``bin`` directory. You could have issues if you build something locally and push those to dotCloud, since those build environments are most likely not the same, and your binary file will not be able to run. 

Helping
=======
All suggestions and pull requests are welcome. If you have an idea or issues, please submit an issue. If you find a bug, or want to contribute, fork the repo, and send a pull request.

Contributors
============
Ken Cochrane < @KenCochrane >
