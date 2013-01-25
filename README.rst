==================
golang-on-dotcloud
==================

Do you have a GO application that you want to run on dotCloud? If so, you have come to the right place. dotCloud doesn't support GO as a native service, but with this Custom service recipe you can still use dotCloud for all of your GO needs.

This is still very early beta so use with caution, and let us know if you have any issues.

Pull requests are welcome.

Version
=======
GO: 1.0.3 

How to use
==========
1. Fork this repo
2. Put your GO source code under src
3. Change the dotcloud.yml
    a. Change the ``build_package`` config variable, to the name of your GO package
    b. Change the ``processes`` so that you can run the correct processes
    c. Add or remove ``ports`` depending on your application needs.
4. Create a dotCloud application
5. Push your code to your new dotCloud application
6. Enjoy!

What to watch out for
=====================
1. Make sure you don't push any of your files from your ``bin`` directory. You could have issues if you build something locally and push those to dotCloud, since those build environments are most likely not the same, and your binary file will not be able to run. 


Contributors
============
Ken Cochrane < @KenCochrane >