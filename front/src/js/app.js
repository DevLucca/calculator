function sendEquation(eq) {
    console.log(eq, {equation:eq})
    $.ajax({
        url: 'http://localhost:1323',
        type: 'post',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            $('#result').val(data.result);
        },
        data: JSON.stringify({ equation: eq })
    });
}

$('.key').on('click', function() {
    let key = $(this).html()
    let result = $('#result').val()
    if (result == '0' && key == '0') {
        return
    } else if (result == '0') {
        result = ''
    }
    switch (key)
    {
        case 'AC':
            $('#result').val('0')
            break;
        case '=':
            sendEquation(result)
            break;
        default:
            $('#result').val(result + key)
    }
})

