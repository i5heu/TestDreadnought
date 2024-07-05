settings = globalsettings;

body = "this is the body yada yada";
data = { body: body };

result = Post("/RAKBtY2U7CrEdANof0fu", data);
ResultIsLikeFile(result.body, "./result.txt");
ThisIsTest();