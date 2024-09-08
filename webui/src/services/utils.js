var utils = {}

utils.since = (timestamp) => {
    const ONE_MINUTE = 60 * 1000;
    const ONE_HOUR = 60 * ONE_MINUTE;
    const ONE_DAY = 24 * ONE_HOUR;
    const ONE_WEEK = 7 * ONE_DAY;

    timestamp = new Date(timestamp);

    const diff = Math.abs(Date.now() - timestamp);
    const elapsedMinutes = Math.floor(diff / ONE_MINUTE);
    const elapsedHours = Math.floor(diff / ONE_HOUR);
    const elapsedDays = Math.floor(diff / ONE_DAY);
    const elapsedWeeks = Math.floor(diff / ONE_WEEK);

    if (elapsedWeeks > 0) {
        return `${elapsedWeeks}w`;
    } else if (elapsedDays > 0) {
        return `${elapsedDays}d`;
    } else if (elapsedHours > 0) {
        return `${elapsedHours}h`;
    } else {
        return `${elapsedMinutes}m`;
    }
}

utils.errorToString = (error) => {
    console.log(error);
    if (error.hasOwnProperty('response')) {
        return error.response.data;1
    } else {
        return error.toString();
    }
}


utils.getCurrentId = () => {
    return localStorage["token"] || null
}

utils.getCurrentUsername = () => {
    return localStorage["username"] || null
}

export default utils