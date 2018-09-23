const assert = require('assert.js');
const lcs = LocalContractStorage;

function assertEqual(func, args, expected, expected_count, msg) {
    const count_of_helper_statement = 46;
    var count = _instruction_counter.count;
    assert.equal(func.apply(null, args), expected);
    assert.equal(_instruction_counter.count - count - count_of_helper_statement, expected_count, msg);
};

// test1.
var test1 = function (k, v) {
    lcs.set(k, v);
    return lcs.get(k);
};

assertEqual(test1, ["k", "1"], "1", 28);
assertEqual(test1, ["k", "12"], "12", 28 + 1);
assertEqual(test1, ["k", "123"], "123", 28 + 2);
assertEqual(test1, ["k1", "1"], "1", 28 + 1);
assertEqual(test1, ["k12", "1"], "1", 28 + 2);

// test2.
var test2 = function (k) {
    lcs.del(k);
};

assertEqual(test2, ["k"], undefined, 12);
assertEqual(test2, ["k1"], undefined, 12);
assertEqual(test2, ["k12"], undefined, 12);
