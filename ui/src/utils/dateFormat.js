import format from 'date-fns/format'

export function dateFormat(time, pattern = 'YYYY-MM-DD HH:mm:ss') {
  return time ? format(time, pattern) : ''
}

export function dateRangeFormat(dateRange, pattern = 'YYYY-MM-DD HH:mm:ss') {
  if (Array.isArray(dateRange)) {
    const dr = dateRange.map(date => dateFormat(date, pattern))

    // startTime and endTime must exist togegher.
    if (dr[0] && dr[1]) {
      return dr
    }
    return []
  }
  return dateRange || []
}

export function todayRangeFormat(date = Date.now()) {
  const start = dateFormat(new Date(date).setHours(0, 0, 0))
  const end = dateFormat(new Date(date).setHours(23, 59, 59))
  return [start, end]
}
