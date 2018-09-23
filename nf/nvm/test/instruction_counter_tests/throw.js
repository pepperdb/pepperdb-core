const assert = require('assert.js');

function assertEqual(func, args, expected, expected_count, msg) {
    const count_of_helper_statement = 46;
    var count = _instruction_counter.count;
    assert.equal(func.apply(null, args), expected);
    assert.equal(_instruction_counter.count - count - count_of_helper_statement, expected_count, msg);
};

// test1
var test1 = function (x) {
    try {
        if (x < 10) {
            return x;
        }
        throw new Error();
    } catch (e) {
        return 0
    };
};
assertEqual(test1, [2], 2, 3);
assertEqual(test1, [10], 0, 17);
