#import <Cocoa/Cocoa.h>
#import <WebKit/Webkit.h>
#import <Security/Security.h>
#import <dispatch/dispatch.h>

typedef void *Menu;
typedef void *MenuItem;
typedef void *Window;

void start();
void attemptQuit();
void mayQuitNow(int quit);
void hideApp();
void hideOtherApps();
void showAllApps();
void invoke(unsigned long id);
void invokeAfter(unsigned long id, long afterNanos);

void setMenuBar(Menu bar);
void setServicesMenu(Menu menu);
void setWindowMenu(Menu menu);
void setHelpMenu(Menu menu);

Menu newMenu(const char *title);
void disposeMenu(Menu menu);
int menuItemCount(Menu menu);
MenuItem menuItem(Menu menu, int index);
void insertMenuItem(Menu menu, MenuItem item, int index);
void removeMenuItem(Menu menu, int index);

MenuItem newMenuItem(const char *title, const char *key, int modifiers);
Menu subMenu(MenuItem item);
void setSubMenu(MenuItem item, Menu subMenu);
MenuItem newMenuSeparator();
void disposeMenuItem(MenuItem item);

Window getKeyWindow();
void bringAllWindowsToFront();
Window newWindow(int styleMask, double x, double y, double width, double height, const char *url);
void closeWindow(Window window);
void setWindowTitle(Window window, const char *title);
void getWindowBounds(Window window, double *x, double *y, double *width, double *height);
void setWindowBounds(Window window, double x, double y, double width, double height);
void bringWindowToFront(Window window);
void minimizeWindow(Window window);
void zoomWindow(Window window);