function main() {
  try {
    const user = getUser();
    // use the "user" here

    // profile has to be here within the scope of user, it cannot be in it's own trycatch block
    const profile = getUserProfile(user.ID);
    // use the "profile" here
  } catch (err) {
    console.log(err);
    // the issue here is that we cannot handle the errors individually they all end up here
  }
}

// this function signature doesn't tell us what is returned, and also that the function can throw
function getUser() {
  // do some get user logic here
  return user;
}
