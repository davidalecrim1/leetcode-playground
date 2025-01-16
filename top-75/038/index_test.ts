import { assertEquals } from "https://deno.land/std@0.224.0/assert/mod.ts";
import { Interval, Solution_Second } from "./index.ts";

Deno.test("Solution Test Case 01", () => {
    const intervals: Interval[] = [
        { start: 0, end: 30 },
        { start: 5, end: 10 },
        { start: 15, end: 20 },
    ];

    const solution = new Solution_Second();
    const result = solution.canAttendMeetings(intervals);

    assertEquals(result, false);
});

Deno.test("Solution Test Case 02", () => {
    const intervals: Interval[] = [
        { start: 5, end: 8 },
        { start: 9, end: 15 },
    ];

    const solution = new Solution_Second();
    const result = solution.canAttendMeetings(intervals);

    assertEquals(result, true);
});
