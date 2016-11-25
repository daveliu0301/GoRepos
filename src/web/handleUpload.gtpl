<html>
<head>
<title></title>
</head>
<body>
<form enctype="multipart/form-data" action="/handleUpload" method="post">
    <input type="file" name="uploadFile"/>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="upload">
</form>
</body>
</html>

