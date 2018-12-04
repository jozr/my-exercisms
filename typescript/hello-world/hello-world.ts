class HelloWorld {
    static hello(name?: string) {
        if (!name) name = "World";
        return `Hello, ${name}!`;
    };
};

export default HelloWorld;
