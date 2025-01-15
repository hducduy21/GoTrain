const config = {
    plugins: ['prettier-plugin-go-template'],
    ovverides: [
        {
            files: ["*.html", "*.tmpl"],
            options: {
                parser: 'go-template',
            },
        },
    ],
}

module.exports = config;
