class Human:
    species = "Homo Sapiens"

    def __init__(self, name, age):
        self.name = name
        self.age = age

    def age_time_10(self):
        return self._age * 10

    @property
    def age(self):
        return self._age

    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, name):
        self._name = name

    @age.setter
    def age(self, age):
        self._age = age

    @classmethod
    def get_species(cls):
        return cls.species

    @staticmethod
    def grunt():
        return "*grunt*"

if __name__ == "__main__":
    human = Human("Murtaza", 18)
    print(human.name)
    print(human.age)
    print(human.get_species())
    print(human.grunt())
