describe('NewDomain', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:first-child')).click();
  });

  it('should have <my-new-domain>', function () {
    var home = element(by.css('my-app my-new-domain'));
    expect(home.isPresent()).toEqual(true);
    expect(home.getText()).toEqual("Home Works!");
  });

});
