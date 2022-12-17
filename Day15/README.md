# Day 15 Notes

## Part 1

For me, this is an exercise about finding an efficient algorithm. Of course there is a brute force method for solving
this exercise, but there's no real learning in that approach and I have a feeling a brute force method would be
very, very slow with the actual data.

Once we know a sensor and it's paired beacon, we can calculate the distance between them. The distance will also
be to maximum x or y distance from the sensor where there are no other beacons. I'll call this "reach".

Now we can figure out if any sensor has reach over the target row. If the sensor's Y position is within reach of the
target row then there is overlap. From the mount of overlap, we can calculate the X positions where there isn't a
sensor.

In the example, our target row is 10.

- Reach = distance between Sensor and Beacon
- Overlaps? = abs(Target Row - Sensor(Y)) <= Reach
- Overlap = abs(abs(Target Row - Sensor(Y)) - Reach)
- X Values = Sensor(X) - Overlap ... Sensor(X) + Overlap

| Sensor | Beacon | Reach | Overlaps? | Overlap | X Value Range |
|--------|--------|-------|-----------|---------|---------------|
| 2, 18  | -2, 15 | 7     | No        |         |               |
| 9, 16  | 10, 16 | 1     | No        |         |               |
| 13, 2  | 15, 3  | 3     | No        |         |               |
| 12, 14 | 10, 16 | 4     | Yes       | 0       | 12            |
| 10, 20 | 10, 16 | 4     | No        |         |               |
| 14, 17 | 10, 16 | 5     | Yes       | 2       | 12:16         |
| 8, 7   | 2, 10  | 9     | Yes       | 6       | 2:14          |
| 2, 0   | 2, 10  | 10    | Yes       | 0       | 2             |
| 0, 11  | 2, 10  | 3     | Yes       | 2       | -2:2          |
| 20, 14 | 25, 17 | 8     | Yes       | 4       | 16:24         |
| 17, 20 | 21, 22 | 6     | Yes       | 4       | 13:21         |
| 16, 7  | 15, 3  | 5     | Yes       | 2       | 14:18         |
| 14, 3  | 15, 3  | 1     | No        |         |               |
| 20, 1  | 15, 3  | 7     | No        |         |               |

Now we know the ranges, but some of them overlap so we should collapse them ranges. The final range is -2:24.

This is 27 positions, but the sensor is at position 2, so we need to remove that - which leaves 26.
