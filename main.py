import PySide2.QtCore
import PySide2.QtGui
import PySide2.QtWidgets


class App(PySide2.QtWidgets.QMainWindow):
    def __init__(self):
        super().__init__(None)


def main():
    app = PySide2.QtWidgets.QApplication()
    gui = App()
    gui.show()
    exit(app.exec_())


if __name__ == '__main__':
    main()
