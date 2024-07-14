Settings = globalSettings;
Settings.baseUrl = "https://example.com";

result = Get("/helloWorld");
console.log("Cache-Control:", result.header["Cache-Control"]);