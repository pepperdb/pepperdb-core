_instruction_counter.incr(1);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 1) throw new Error('_instruction_counter.count error, expected ' + 1 + ', actual is ' + _instruction_counter.count);

_instruction_counter.incr(2);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 3) throw new Error('_instruction_counter.count error, expected ' + 3 + ', actual is ' + _instruction_counter.count);

_instruction_counter.incr(3);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 6) throw new Error('_instruction_counter.count error, expected ' + 6 + ', actual is ' + _instruction_counter.count);


_instruction_counter.count = 0123;
if (_instruction_counter.count != 6) throw new Error('_instruction_counter.count error, expected ' + 6 + ', actual is ' + _instruction_counter.count);

_instruction_counter.incr(4);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 10) throw new Error('_instruction_counter.count error, expected ' + 10 + ', actual is ' + _instruction_counter.count);

delete _instruction_counter.count;
if (_instruction_counter.count != 10) throw new Error('_instruction_counter.count error, expected ' + 10 + ', actual is ' + _instruction_counter.count);

_instruction_counter.incr(5);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 15) throw new Error('_instruction_counter.count error, expected ' + 15 + ', actual is ' + _instruction_counter.count);

_instruction_counter.incr(-1);
console.log('count = ' + _instruction_counter.count);
if (_instruction_counter.count != 15) throw new Error('_instruction_counter.count error, expected ' + 15 + ', actual is ' + _instruction_counter.count);
