Settings = globalSettings;

body = "this is the body yada yada";
data = { body: body };

result = Post("/RAKBtY2U7CrEdANof0fu", data);
console.log('after Post ->');
ResultIsLikeFile(result.body, "./result.txt");
console.log('result.body ->', result.body);