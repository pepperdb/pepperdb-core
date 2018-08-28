interface Console {
    debug(message?: any, ...optionalMessage: any[]): void;
    warn(message?: any, ...optionalMessage: any[]): void;
    info(message?: any, ...optionalMessage: any[]): void;
    log(message?: any, ...optionalMessage: any[]): void;
    error(message?: any, ...optionalMessage: any[]): void;
}

