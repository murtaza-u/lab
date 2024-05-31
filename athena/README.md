# Distance between 2 points on a sphere

## Great Circle Distance Approach

* Let `P1(x1, y1)` and `P2(x2, y2)` be 2 points on a sphere of radius R
* `P1` is the location of the user
* `P2` is the location of the object
* R is the radius of the earth (i.e. 6372 km)
* Distance b/w `P1` and `P2` is given by,

```
D = acos(sin(x1) * sin(x2) + cos(x1) * cos(x2) * cos(y1 - y2)) * R
```

## Athena SQL Query

* User's coordinates: (0.7381, 1.6443)

```sql
SELECT x, y
FROM heatmaps.india
WHERE acos(sin(x) * sin(0.7381) + cos(x) * cos(0.7381) * cos(y - 1.6443)) * 6371 <= 5
LIMIT 10;
```

## References

* <http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates>
* <http://en.wikipedia.org/wiki/Great-circle_distance>
