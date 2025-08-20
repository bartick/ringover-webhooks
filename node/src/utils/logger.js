// Logger configuration
// logger options error, warn, info, http, verbose, debug, silly

const { createLogger, format, transports } = require('winston');

const {
  combine,
  timestamp,
  printf,
  prettyPrint,
  json,
  colorize,
  align,
} = format;

require('winston-daily-rotate-file');

// Utils
const { NODE_ENV } = require('./config');

const myformat = combine(
    colorize(),
    timestamp(),
    align(),
    printf((info) => {
        let message = `${info.timestamp} ${info.level}: ${info.message}`;
        if (info.stack) message += `\nTraceback: ${info.stack}`;
        return message;
    })
);

const ServerTransport = new transports.DailyRotateFile({
    filename: './logs/server-%DATE%.log',
    datePattern: 'YYYY-MM-DD',
    zippedArchive: true,
    maxSize: '20m',
    maxFiles: '7d',
    level: 'info',
    format: combine(json(), timestamp(), prettyPrint()),
});

const ErrorTransport = new transports.DailyRotateFile({
    filename: './logs/error-%DATE%.log',
    datePattern: 'YYYY-MM-DD',
    zippedArchive: true,
    maxSize: '20m',
    maxFiles: '7d',
    format: combine(json(), timestamp(), prettyPrint()),
    level: 'error',
});

const DebugTransport = new transports.DailyRotateFile({
    filename: './logs/debug-%DATE%.log',
    datePattern: 'YYYY-MM-DD',
    zippedArchive: true,
    maxSize: '20m',
    maxFiles: '7d',
    format: combine(json(), timestamp(), prettyPrint()),
    level: 'debug',
});

const logger = new createLogger({
    transports: [
        ServerTransport,
        ErrorTransport,
        DebugTransport,
        new transports.Console({
            level: NODE_ENV === 'production' ? 'info' : 'debug',
            handleExceptions: true,
            format: myformat,
        }),
    ],
    exitOnError: false,
});

logger.stream = {
    write(message) {
        logger.debug(message);
    },
};

module.exports = logger;