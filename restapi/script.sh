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


   *)
     echo 'other OS'
     ;;
esac



#notizen
#datenbank erstellen, query ausführen, per xo models erstellen, backend starten
