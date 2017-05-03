describe('Settings', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:last-child')).click();
  });

  it('should have <my-settings>', function () {
    var home = element(by.css('my-app my-settings'));
    expect(home.isPresent()).toEqual(true);
  });

});
