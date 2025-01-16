class Interval {
    start: number;
    end: number;

    constructor(start: number, end: number) {
        this.start = start;
        this.end = end;
    }
}

// Time Complexity: O(n) (actually n/2 but Big O doesn't care about it.)
// This solution only works if the meetings are ordered,
// and I believe that should be considered
// hence it's not the best solution.
class Solution_First {
    canAttendMeetings(intervals: Interval[]): boolean {
        if (intervals.length == 1) {
            return true;
        }

        for (let i = 1; i < intervals.length; i++) {
            if (intervals[i - 1].end <= intervals[i].start) {
                continue;
            } else {
                return false;
            }
        }

        return true;
    }
}

// Time Complexity: O(n log n)
// O(n log n) for sorting algorithm
// This is the best solution considering sorting.
class Solution_Second {
    canAttendMeetings(intervals: Interval[]): boolean {
        if (intervals.length == 1) {
            return true;
        }

        intervals.sort((a, b) => {
            return a.start - b.start;
        });

        for (let i = 1; i < intervals.length; i++) {
            if (intervals[i - 1].end <= intervals[i].start) {
                continue;
            } else {
                return false;
            }
        }

        return true;
    }
}

export { Interval, Solution_First, Solution_Second };
