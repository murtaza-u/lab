const name = "murtaza"

if name == "murtaza":
  echo "Genius"
elif name == "mark":
  echo "Smart"
else:
  echo "Dumb"

case name
  of "murtaza", "Murtaza":
    echo "Genius"
  of "mark":
    echo "Smart"
  else:
    echo "Dumb"

from std/strutils import parseInt

echo "Enter a number:"
let x = parseInt(readLine(stdin))
case x
  of 0..10:
    echo "[0, 19]"
  else:
    discard

echo "Enter a number b/w 0 to 10:"
var n = parseInt(readline(stdin))
while n < 0 or n > 10:
  echo "Please enter a number b/w 0 to 10 only:"
  n = parseInt(readline(stdin))

for i in countup(1, 10):
  echo i

for i in 1..10:
  echo i

var i = 1
while i <= 10:
  echo i
  inc i
