import unittest


class Solution(object):
    @staticmethod
    def calorie_counting_part_1(content):
        """
        :type content: str
        :rtype: int
        """
        max_calo = -1  # Max calo carried by an elf
        calo_so_far = 0  # Total calo carried by an elf so far
        for line in content.splitlines():
            if line.strip():
                calo = int(line)
                calo_so_far += calo
            else:
                # Line is empty -> end of inventory for current elf
                # Check if total calo exceeds max before reset
                if calo_so_far > max_calo:
                    max_calo = calo_so_far
                calo_so_far = 0
        return max_calo

    @staticmethod
    def calorie_counting_part_2(content):
        """
        :type content: str
        :rtype: int
        """
        max_calo_1 = -1  # Top 1 calo carried by an elf
        max_calo_2 = -1  # Top 2 calo carried by an elf
        max_calo_3 = -1  # Top 3 calo carried by an elf
        calo_so_far = 0  # Total calo carried by an elf so far
        for line in content.splitlines():
            if line.strip():
                calo = int(line)
                calo_so_far += calo
            else:
                # Line is empty -> end of inventory for current elf
                # Check if total calo exceeds each max, swap their values if so
                # This is because we want to persist the values of each max as
                # they are likely to be the candidate for next max
                if calo_so_far > max_calo_1:
                    calo_so_far, max_calo_1 = max_calo_1, calo_so_far
                if calo_so_far > max_calo_2:
                    calo_so_far, max_calo_2 = max_calo_2, calo_so_far
                if calo_so_far > max_calo_3:
                    calo_so_far, max_calo_3 = max_calo_3, calo_so_far
                # Reset total calo
                calo_so_far = 0
        return max_calo_1 + max_calo_2 + max_calo_3


class CalorieCountingTestCase(unittest.TestCase):
    def setUp(self):
        with open('input.txt') as file:
            self.content = file.read()

    def tearDown(self):
        pass

    def test_part_1(self):
        s = Solution()
        self.assertEqual(s.calorie_counting_part_1(self.content), 66487)

    def test_part_2(self):
        s = Solution()
        self.assertEqual(s.calorie_counting_part_2(self.content), 197301)


if __name__ == '__main__':
    unittest.main()
