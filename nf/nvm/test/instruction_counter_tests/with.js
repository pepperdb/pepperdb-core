const assert = require('assert.js');

function assertEqual(func, args, expected, expected_count, msg) {
    const count_of_helper_statement = 46;
    var count = _instruction_counter.count;
    assert.equal(func.apply(null, args), expected);
    assert.equal(_instruction_counter.count - count - count_of_helper_statement, expected_count, msg);
};

// test1.
var test1 = function (x) {
    with(x)
    return a + b;
};
assertEqual(test1, [{
    a: 1,
    b: 2
}], 3, 3);

// test2
var test2 = function (x) {
    with(x[0]) {
        return a + b;
    }
};
assertEqual(test2, [
    [{
        a: 1,
        b: 2
    }]
], 3, 7);

// test3
var gen3X = function (a, b) {
    return {
        val: function () {
            return {
                a: a,
                b: b
            };
        }
    }
};
var test3 = function (x) {
    with(x.val()) {
        if (a < b) {
            return a + b;
        } else if (a > b) {
            return a * 2 + b;
        } else {
            return 0;
        }
    }
};
assertEqual(test3, [gen3X(3, 5)], 8, 18);
assertEqual(test3, [gen3X(5, 3)], 13, 24);
assertEqual(test3, [gen3X(4, 4)], 0, 18);
