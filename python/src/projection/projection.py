from abc import ABC, abstractmethod


class Projection(ABC):

    @abstractmethod
    def accept(self, event):
        pass
