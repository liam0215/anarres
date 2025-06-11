# ROOT_DIR=..
CC = gcc
AR = ar
CFLAGS += -g
CFLAGS += -I/data/liam/anarres/deps/qpl/install/include
# All warnings are errors
CFLAGS += -Werror

# Enable optimizations
CFLAGS += -O3
# Disable asserts etc
CFLAGS += -DNDEBUG
# Generate optimized code
CFLAGS += -march=native -flto=auto -ffast-math
CFLAGS += -MMD
# LDFLAGS += -L $(ROOT_DIR)/deps/qpl/install/lib
# LIBS += -l:libqpl.a

TARGET = libqpl_shim.a
MAIN_SRC = qpl_shim.c
# ADDITIONAL_SRCS =
MAIN_OBJ = $(MAIN_SRC:.c=.o)
# ADDITIONAL_OBJS = $(ADDITIONAL_SRCS:.cc=.o)

DEP = $(MAIN_OBJ:.o=.d)
# DEP += $(ADDITIONAL_OBJS:.o=.d)

all: $(TARGET)

ifneq ($(MAKECMDGOALS),clean)
-include $(DEP)   # include all dep files in the makefile
endif

$(TARGET): $(MAIN_OBJ)
	$(AR) rcs $@ $^

%.o: %.c
	$(CC) -c $(CFLAGS) $< -o $@

clean:
	rm -f $(MAIN_OBJ) $(TARGET) $(DEP)

.PHONY: all clean

# handy for debugging
print-%  : ; @echo $* = $($*)
