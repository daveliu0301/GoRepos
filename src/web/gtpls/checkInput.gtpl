<html>
<head>
<title></title>
</head>
<body>
<form action="/checkInput" method="post">
    required:<input type="text" name="required">
    number:<input type="text" name="number">
    chinese:<input type="text" name="chinese">
    english:<input type="text" name="english">

    <select name="fruit">
    <option value="apple">apple</option>
    <option value="pear">pear</option>
    <option value="banana">banana</option>
    <option value="orange">orange</option>
    </select>

    <input type="radio" name="gender" value="1">male
    <input type="radio" name="gender" value="2">female
    <input type="radio" name="gender" value="11">unkown

    <input type="checkbox" name="interest" value="football">football
    <input type="checkbox" name="interest" value="basketball">basketball
    <input type="checkbox" name="interest" value="tennis">tennis

    password:<input type="password" name="password">
    <input type="submit" value="submit">
</form>
</body>
</html>

