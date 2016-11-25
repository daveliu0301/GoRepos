<html>
<head>
<title></title>
</head>
<body>
<form action="/preventRepeat" method="post">
    input:<input type="text" name="input">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="submit">
</form>
</body>
</html>

