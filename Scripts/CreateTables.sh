#!/bin/sh

case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
     ;;

   Linux)
     echo 'Linux'
     ;;

   CYGWIN*|MINGW32*|MSYS*)
     echo 'MS Windows'
     ;;

   # Add here more strings to compare
   # See correspondence table at the bottom of this answer

   *)
     echo 'other OS'
     ;;
esac
