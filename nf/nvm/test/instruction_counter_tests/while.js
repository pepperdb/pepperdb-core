const assert = require('assert.js');

function assertEqual(func, args, expected, expected_count, msg) {
    const count_of_helper_statement = 46;
    var count = _instruction_counter.count;
    assert.equal(func.apply(null, args), expected);
    assert.equal(_instruction_counter.count - count - count_of_helper_statement, expected_count, msg);
};

// test1.
var test1 = function (x) {
    var ret = 1;
    var i = x;
    while (ret < 1024 && i > 0) {
        ret *= 2;
        i--;
    }
    return ret;
};
assertEqual(test1, [0], 1, 9);
assertEqual(test1, [2], 4, 15 * 2 + 9);
assertEqual(test1, [10], 1024, 15 * 10 + 9);
assertEqual(test1, [11], 1024, 15 * 10 + 9);

// test2.
var test2 = function (x) {
    var ret = 1;
    var i = x;
    do {
        ret *= 2;
        i--;
    } while (ret < 1024 && i > 0);
    return ret;
}
assertEqual(test2, [0], 2, 15);
assertEqual(test2, [2], 4, 15 * 2);
assertEqual(test2, [10], 1024, 15 * 10);
assertEqual(test2, [11], 1024, 15 * 10);

// test3.
var test3 = function (x) {
    var ret = 1;
    while (ret < 1024) ret *= 2;
    return ret;
};
assertEqual(test3, [1], 1024, 63);
assertEqual(test3, [2], 1024, 63);
assertEqual(test3, [10], 1024, 63);
