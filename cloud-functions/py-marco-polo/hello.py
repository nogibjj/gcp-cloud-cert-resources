def hello_world(request):
    """
    Responds to an HTTP request with a message based on the "name" parameter.
    """
    request_json = request.get_json()

    if request_json and "name" in request_json:
        name = request_json["name"]

        if name.lower() == "marco":
            return "Polo"
        else:
            return "I don't know you"

    return "Please provide a 'name' parameter in the request body."
