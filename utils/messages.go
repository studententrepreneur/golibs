package utils


var Errors = errorMessages()

type errorMessagesRegistry struct {
    Created   string
    Accepted string
    NoContent  string
	MovedPermanently  string
	Found  string
	BadRequest  string
	Unauthorized  string
	Forbidden  string
	NotFound  string
	InternalServerError  string
}

func errorMessages() *errorMessagesRegistry {
    return &errorMessagesRegistry{
        Created:   "Success! Your new resource has been created.",
        Accepted: "Thanks for your request. We're processing it now and will let you know when it's done.",
        NoContent:  "Request processed successfully. No content to display at this time.",
		MovedPermanently:  "Oops! The resource you're looking for has moved to a new location. Don't worry, we've redirected you.",
		Found:  "Oops! The resource you're looking for has moved to a new location. Don't worry, we've redirected you.",
		BadRequest:  "Uh-oh! It looks like something went wrong with your request. Please check your inputs and try again.",
		Unauthorized:  "This action requires authentication. Please sign in to continue.",
		Forbidden:  "This action requires authentication. Please sign in to continue.",
		NotFound:  "Oops! We couldn't find the page you're looking for. Please check the URL and try again.",
		InternalServerError:  "Sorry, something went wrong on our end. Please try again later or contact support for assistance.",
    }
}
