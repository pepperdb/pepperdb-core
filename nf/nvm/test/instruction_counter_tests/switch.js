const assert = require('assert.js');

function assertEqual(func, args, expected, expected_count, msg) {
    const count_of_helper_statement = 46;
    var count = _instruction_counter.count;
    assert.equal(func.apply(null, args), expected);
    assert.equal(_instruction_counter.count - count - count_of_helper_statement, expected_count, msg);
};


// test1.
var gen1X = function (a) {
    var o = {};
    o.get = function () {
        return a;
    }
    return o;
}
var test1 = function (x) {
    var ret = 0;
    switch (x.get()) {
        case 1:
            ret = 1 + 2;
            break;
        case 2:
            ret = 2 * 2;
            break;
        case 3:
            ret = 3 * 3;
        case 4:
            ret += 4 * 4;
            break;
        case 100:
            return 100;
        case 101:
            return 3 * x.get();
        default:
            ret = x.get() * 2 + 13;
    }
    return ret;
};
assertEqual(test1, [gen1X(1)], 3, 18);
assertEqual(test1, [gen1X(2)], 4, 18);
assertEqual(test1, [gen1X(3)], 25, 24);
assertEqual(test1, [gen1X(4)], 16, 18);
assertEqual(test1, [gen1X(100)], 100, 12);
assertEqual(test1, [gen1X(101)], 303, 27);
assertEqual(test1, [gen1X(50)], 113, 33);

// test2.
var gen2X = function (a) {
    return [a];
}
var test2 = function (x) {
    var ret = 0;
    switch (x[0]) {
        case 1:
            ret = 1 + 2;
            break;
        case 2:
            ret = 2 * 2;
            break;
        case 3:
            ret = 3 * 3;
        case 4:
            ret += 4 * 4;
            break;
        case 100:
            return 100;
        case 101:
            return 3 * x[0];
        default:
            ret = x[0] * 2 + 13;
    }
    return ret;
};
assertEqual(test2, [gen2X(1)], 3, 10);
assertEqual(test2, [gen2X(2)], 4, 10);
assertEqual(test2, [gen2X(3)], 25, 16);
assertEqual(test2, [gen2X(4)], 16, 10);
assertEqual(test2, [gen2X(100)], 100, 4);
assertEqual(test2, [gen2X(101)], 303, 11);
assertEqual(test2, [gen2X(50)], 113, 17);
