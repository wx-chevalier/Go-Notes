/* disable eslint */
// http://thecodebarbarian.com/algorithm-interview-questions-in-js-glob-matching.html

function patternMatchesRecursive(pattern, str) {
  if (!pattern.includes('*')) {
    // No wildcards, so must be an exact match
    return pattern === str;
  }

  const starIndex = pattern.indexOf('*');
  const leftPattern = pattern.substr(0, starIndex);
  const rightPattern = pattern.substr(starIndex + 1);

  if (leftPattern !== str.substr(0, starIndex)) {
    // Non-wildcard characters at the start of `pattern` are different from
    // the start of `str`, for example `ab*` and `ba`
    return false;
  }

  if (rightPattern.length === 0) {
    // Nothing left in pattern
    return str.startsWith(leftPattern);
  }

  for (let numChars = 0; numChars < str.length - starIndex; ++numChars) {
    // Check to see if the remaining part of `pattern` matches some part of `str`
    if (
      patternMatchesRecursive(rightPattern, str.substr(starIndex + numChars))
    ) {
      return true;
    }
  }

  return false;
}

function patternMatchesDynamic(pattern, str) {
  if (!pattern.includes('*')) {
    // No wildcards, so must be an exact match
    return pattern === str;
  }

  // 备忘录
  const arr = [];
  for (let i = 0; i <= pattern.length; ++i) {
    arr.push([]);

    // 将所有的初始化值设为 false
    for (let j = 0; j <= str.length; ++j) {
      arr[i].push(false);
    }
  }

  // Empty pattern matches empty string
  arr[0][0] = true;

  // Empty str only matches if pattern is '*'
  for (let i = 1; i < pattern.length; ++i) {
    arr[i][0] = pattern === '*';
  }

  // Empty pattern is always false

  // Build up array using recurrence relationship
  for (let i = 1; i <= pattern.length; ++i) {
    for (let j = 1; j <= str.length; ++j) {
      if (pattern[i - 1] === '*') {
        // Two cases: either we use the wildcard, in which case `arr[i][j - 1]` must be true for a match,
        // or we don't, in which case `arr[i - 1][j]` must be true
        arr[i][j] = arr[i - 1][j] || arr[i][j - 1];
      } else {
        // If no wildcard, chars must be equal and previous substrs must match
        arr[i][j] = pattern[i - 1] === str[j - 1] && arr[i - 1][j - 1];
      }
    }
  }

  return arr[pattern.length][str.length];
}

function patternMatches(pattern, str, option = 'dynamic') {
  if (option === 'dynamic') {
    return patternMatchesDynamic(pattern, str);
  }

  return patternMatchesRecursive(pattern, str);
}

module.exports = patternMatches;
