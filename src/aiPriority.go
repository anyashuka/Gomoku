package gomoku

// const capture10 = maxInt - 100000
// const break5Align = maxInt - 200000
// const willBeCaptured8 = -(maxInt - 300000)
// const blockWin = maxInt - 400000
// // const block4 = maxInt - 500000
// const willBeCaptured2 = -42e14
// const align5Win = (maxInt - 1000)
// const alignFree4 = 42e14
// const blockFree3 = 42e11 + 500
// const capture2 = 42e11
// const align4 = 42e10 + 200
// const alignFree3 = 42e10

const capture10 = maxInt - 100000
const break5Align = maxInt - 200000
const willBeCaptured8 = -(maxInt - 300000)
const align5Win = maxInt - 400000
const blockWin = maxInt - 500000

// const block4 = blockWin // rm !!! redundant
const willBeCaptured2 = -42e14
const alignFree4 = 42e13
const blockFree3 = 42e12
const capture2 = 42e11
const alignFree3 = 42e10
const block2 = 42e7

// const align4 = 42e10 + 200 // rm !!! don't need if flanked, no point

// const block2 =

// const align3 = 42e9
// 1. Capture 10
// 2. willBeCaptured (if capture 8)
// 3. Block align 5 ( includes block 4)
// 4. willBeCaptured
// 5. Align 5
// 6. Free 4
// 7. Block Free 3
// 8. Capture
// 9. Free 3
// 10. Block 2
// 11. Free 2

//  maxINT = 9223372036854775807

// ISSUES:
// - I still don’t think check flanked is doing what it needs to to
// differentiate free3 from not free3
// - captureAttackDefend: blockFree3 trumps align5Win
