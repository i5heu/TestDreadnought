Settings = globalSettings;

body = "this is the body yada yada";
data = { body: body };

result = Post("/RAKBtY2U7CrEdANof0fu", data);
console.log('after Post ->');
ResultIsLikeGlobalFile(result.body, "./GlobalResultExample.txt");
Log('result.body ->', result.body);