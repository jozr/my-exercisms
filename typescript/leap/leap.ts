const isLeapYear = (year: number) => (year % 400 === 0 || year % 4 === 0 && year % 100 !== 0)

export default isLeapYear
