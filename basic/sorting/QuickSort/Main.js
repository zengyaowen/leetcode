var buf = '';

process.stdin.on('readable', function () {
    var chunk = process.stdin.read();
    if (chunk) buf += chunk.toString();
});

let getInputArgs = line => {
    return line.split(' ').filter(s => s !== '').map(x => parseInt(x));
}

function quickSort(nums, left, high) {
    if (left >= high) {
        return;
    }
    
    let i = left - 1;
    let j = high + 1;
    let x = nums[(left + high) >> 1];
    while (i < j) {
        while (nums[++i] < x);
        while (nums[--j] > x);
        if (i < j) {
            const t = nums[i];
            nums[i] = nums[j];
            nums[j] = t;
        }
    }
    quickSort(nums, left, j);
    quickSort(nums, j + 1, high);
}



process.stdin.on('end', function () {
    buf.split('\n').forEach(function (line, lineIdx) {
        if (lineIdx % 2 === 1) {
            nums = getInputArgs(line);
            quickSort(nums, 0, nums.length - 1);
            console.log(nums.join(' '));
        }

    });
});